module Test.Streams where

import Prelude

import Control.Alt ((<|>))
import Data.Array (replicate)
import Data.Either (Either(..))
import Data.Foldable (for_)
import Data.String (joinWith)
import Effect (Effect)
import Effect.Aff (Aff, Milliseconds(..), attempt, delay, makeAff, nonCanceler, parallel, runAff_, sequential, throwError)
import Effect.Class (liftEffect)
import Effect.Console (log)
import Effect.Console as Console
import Effect.Exception (error, message)
import Node.Encoding (Encoding(..))
import Node.EventEmitter (once_)
import Node.FS.Stream (createWriteStream, createReadStream)
import Node.FS.Perms (permsAll)
import Node.FS.Sync as Sync
import Node.Process as Process
import Node.Stream as Stream

main :: Effect Unit
main = runAff_ (case _ of
  Left err -> do
    Console.error (message err)
    Process.exit' 1
  Right _ -> pure unit) do
  liftEffect $ Sync.mkdir' "tmp/pipe-regression" { recursive: true, mode: permsAll }
  for_
    [ { name: "multichunk", content: joinWith "" (replicate 4096 "arbitrary pipe data: αβγ\n") }
    , { name: "empty", content: "" }
    ] \test -> do
      let source = "tmp/pipe-regression/" <> test.name <> "-source.txt"
      let destination = "tmp/pipe-regression/" <> test.name <> "-destination.txt"
      liftEffect do
        Sync.writeTextFile UTF8 source test.content
        Sync.writeTextFile UTF8 destination "old destination must be replaced"
      result <- sequential
        ( parallel (attempt $ copyFile source destination)
            <|> parallel (delay (Milliseconds 5000.0) $> Left (error $ "pipe " <> test.name <> " timed out"))
        )
      case result of
        Left err -> throwError err
        Right _ -> pure unit
      actual <- liftEffect $ Sync.readTextFile UTF8 destination
      if actual /= test.content then
        throwError $ error $ "pipe " <> test.name <> " produced the wrong content"
      else liftEffect $ log $ "PASS file pipe " <> test.name

copyFile :: String -> String -> Aff Unit
copyFile source destination = makeAff \done -> do
  r <- createReadStream source
  w <- createWriteStream destination
  r # once_ Stream.errorH (done <<< Left)
  w # once_ Stream.errorH (done <<< Left)
  w # once_ Stream.finishH (done $ Right unit)
  Stream.pipe r w
  pure nonCanceler
