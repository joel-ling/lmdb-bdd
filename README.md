```bash
$ go test
```
```gherkin
Feature: LMDB
  In order to verify the effects of software interacting with LMDB
  As a software engineer practising behaviour-driven development
  I need to establish some baseline observations

  Scenario: Get non-existent record                                           # features/lmdb.feature:6
    Given there is a new LMDB environment "env"                               # new-lmdb-env.go:19 -> github.com/joel-ling/lmdb-bdd/test.newLMDBEnv
    When I get a record "key" in DB "deebee" of environment "env"             # get-record.go:18 -> github.com/joel-ling/lmdb-bdd/test.getRecord
    Then I should see an error "MDB_NOTFOUND"                                 # see-error.go:19 -> github.com/joel-ling/lmdb-bdd/test.seeError
    And the ID of last committed transaction in environment "env" should be 0 # id-last-commit.go:21 -> github.com/joel-ling/lmdb-bdd/test.idLastCommit

  Scenario: Put record                                                         # features/lmdb.feature:12
    Given there is a new LMDB environment "env"                                # new-lmdb-env.go:19 -> github.com/joel-ling/lmdb-bdd/test.newLMDBEnv
    When I put a record "key" "value" in DB "deebee" of environment "env"      # put-record.go:25 -> github.com/joel-ling/lmdb-bdd/test.putRecord
    Then the ID of last committed transaction in environment "env" should be 1 # id-last-commit.go:21 -> github.com/joel-ling/lmdb-bdd/test.idLastCommit

  Scenario: Put and then get record                                           # features/lmdb.feature:17
    Given there is a new LMDB environment "env"                               # new-lmdb-env.go:19 -> github.com/joel-ling/lmdb-bdd/test.newLMDBEnv
    When I put a record "key" "value" in DB "deebee" of environment "env"     # put-record.go:25 -> github.com/joel-ling/lmdb-bdd/test.putRecord
    And I get a record "key" in DB "deebee" of environment "env"              # get-record.go:18 -> github.com/joel-ling/lmdb-bdd/test.getRecord
    Then I should see a value "value"                                         # see-value.go:18 -> github.com/joel-ling/lmdb-bdd/test.seeValue
    And I should see no error                                                 # see-no-error.go:18 -> github.com/joel-ling/lmdb-bdd/test.seeNoError
    And the ID of last committed transaction in environment "env" should be 1 # id-last-commit.go:21 -> github.com/joel-ling/lmdb-bdd/test.idLastCommit

  Scenario: Get record put in different environment                         # features/lmdb.feature:25
    Given there is a new LMDB environment "primary"                         # new-lmdb-env.go:19 -> github.com/joel-ling/lmdb-bdd/test.newLMDBEnv
    And there is a new LMDB environment "secondary"                         # new-lmdb-env.go:19 -> github.com/joel-ling/lmdb-bdd/test.newLMDBEnv
    And there is a record "key" "value" in DB "DB" of environment "primary" # put-record.go:25 -> github.com/joel-ling/lmdb-bdd/test.putRecord
    When I get a record "key" in DB "DB" of environment "secondary"         # get-record.go:18 -> github.com/joel-ling/lmdb-bdd/test.getRecord
    Then I should see an error "MDB_NOTFOUND"                               # see-error.go:19 -> github.com/joel-ling/lmdb-bdd/test.seeError

  Scenario: Get records using a cursor                                # features/lmdb.feature:32
    Given there is a new LMDB environment "env"                       # new-lmdb-env.go:19 -> github.com/joel-ling/lmdb-bdd/test.newLMDBEnv
    And there is a record "41" "A" in DB "ASCII" of environment "env" # put-record.go:25 -> github.com/joel-ling/lmdb-bdd/test.putRecord
    And there is a record "43" "C" in DB "ASCII" of environment "env" # put-record.go:25 -> github.com/joel-ling/lmdb-bdd/test.putRecord
    And there is a record "42" "B" in DB "ASCII" of environment "env" # put-record.go:25 -> github.com/joel-ling/lmdb-bdd/test.putRecord
    When I open a cursor to DB "ASCII" of environment "env"           # open-cursor.go:18 -> github.com/joel-ling/lmdb-bdd/test.openCursor
    And I get the next record using the cursor                        # get-next-record.go:18 -> github.com/joel-ling/lmdb-bdd/test.getNextRecord
    Then I should see a record "41" "A"                               # see-record.go:18 -> github.com/joel-ling/lmdb-bdd/test.seeRecord
    And I should see no error                                         # see-no-error.go:18 -> github.com/joel-ling/lmdb-bdd/test.seeNoError
    When I get the next record using the cursor                       # get-next-record.go:18 -> github.com/joel-ling/lmdb-bdd/test.getNextRecord
    Then I should see a record "42" "B"                               # see-record.go:18 -> github.com/joel-ling/lmdb-bdd/test.seeRecord
    And I should see no error                                         # see-no-error.go:18 -> github.com/joel-ling/lmdb-bdd/test.seeNoError
    When I get the next record using the cursor                       # get-next-record.go:18 -> github.com/joel-ling/lmdb-bdd/test.getNextRecord
    Then I should see a record "43" "C"                               # see-record.go:18 -> github.com/joel-ling/lmdb-bdd/test.seeRecord
    And I should see no error                                         # see-no-error.go:18 -> github.com/joel-ling/lmdb-bdd/test.seeNoError
    When I get the next record using the cursor                       # get-next-record.go:18 -> github.com/joel-ling/lmdb-bdd/test.getNextRecord
    Then I should see an error "MDB_NOTFOUND"                         # see-error.go:19 -> github.com/joel-ling/lmdb-bdd/test.seeError

  Scenario: Get records using a cursor, resetting and renewing read transaction # features/lmdb.feature:50
    Given there is a new LMDB environment "env"                                 # new-lmdb-env.go:19 -> github.com/joel-ling/lmdb-bdd/test.newLMDBEnv
    And there is a record "41" "A" in DB "ASCII" of environment "env"           # put-record.go:25 -> github.com/joel-ling/lmdb-bdd/test.putRecord
    And there is a record "42" "B" in DB "ASCII" of environment "env"           # put-record.go:25 -> github.com/joel-ling/lmdb-bdd/test.putRecord
    And there is a record "43" "C" in DB "ASCII" of environment "env"           # put-record.go:25 -> github.com/joel-ling/lmdb-bdd/test.putRecord
    When I open a cursor to DB "ASCII" of environment "env"                     # open-cursor.go:18 -> github.com/joel-ling/lmdb-bdd/test.openCursor
    And I get the next record using the cursor                                  # get-next-record.go:18 -> github.com/joel-ling/lmdb-bdd/test.getNextRecord
    Then I should see a record "41" "A"                                         # see-record.go:18 -> github.com/joel-ling/lmdb-bdd/test.seeRecord
    And I should see no error                                                   # see-no-error.go:18 -> github.com/joel-ling/lmdb-bdd/test.seeNoError
    When I get the next record using the cursor                                 # get-next-record.go:18 -> github.com/joel-ling/lmdb-bdd/test.getNextRecord
    Then I should see a record "42" "B"                                         # see-record.go:18 -> github.com/joel-ling/lmdb-bdd/test.seeRecord
    And I should see no error                                                   # see-no-error.go:18 -> github.com/joel-ling/lmdb-bdd/test.seeNoError
    When I reset the transaction relating to the cursor                         # reset-cursor.go:18 -> github.com/joel-ling/lmdb-bdd/test.resetCursor
    And I renew the transaction relating to the cursor, and the cursor          # renew-cursor.go:18 -> github.com/joel-ling/lmdb-bdd/test.renewCursor
    And I get the next record using the cursor, specifying the previous key     # get-next-record-set.go:19 -> github.com/joel-ling/lmdb-bdd/test.getNextRecordSet
    Then I should see a record "43" "C"                                         # see-record.go:18 -> github.com/joel-ling/lmdb-bdd/test.seeRecord
    And I should see no error                                                   # see-no-error.go:18 -> github.com/joel-ling/lmdb-bdd/test.seeNoError

  Scenario: List databases in environment                                # features/lmdb.feature:68
    Given there is a new LMDB environment "env"                          # new-lmdb-env.go:19 -> github.com/joel-ling/lmdb-bdd/test.newLMDBEnv
    And there is a record "key" "value" in DB "db0" of environment "env" # put-record.go:25 -> github.com/joel-ling/lmdb-bdd/test.putRecord
    And there is a record "key" "value" in DB "db1" of environment "env" # put-record.go:25 -> github.com/joel-ling/lmdb-bdd/test.putRecord
    And there is a record "key" "value" in DB "db2" of environment "env" # put-record.go:25 -> github.com/joel-ling/lmdb-bdd/test.putRecord
    When I open a cursor to the root database of environment "env"       # open-cursor-root.go:18 -> github.com/joel-ling/lmdb-bdd/test.openCursorRoot
    And I get the next record using the cursor                           # get-next-record.go:18 -> github.com/joel-ling/lmdb-bdd/test.getNextRecord
    Then I should see a key "db0"                                        # see-key.go:18 -> github.com/joel-ling/lmdb-bdd/test.seeKey
    And I should see no error                                            # see-no-error.go:18 -> github.com/joel-ling/lmdb-bdd/test.seeNoError
    When I get the next record using the cursor                          # get-next-record.go:18 -> github.com/joel-ling/lmdb-bdd/test.getNextRecord
    Then I should see a key "db1"                                        # see-key.go:18 -> github.com/joel-ling/lmdb-bdd/test.seeKey
    And I should see no error                                            # see-no-error.go:18 -> github.com/joel-ling/lmdb-bdd/test.seeNoError
    When I get the next record using the cursor                          # get-next-record.go:18 -> github.com/joel-ling/lmdb-bdd/test.getNextRecord
    Then I should see a key "db2"                                        # see-key.go:18 -> github.com/joel-ling/lmdb-bdd/test.seeKey
    And I should see no error                                            # see-no-error.go:18 -> github.com/joel-ling/lmdb-bdd/test.seeNoError
    When I get the next record using the cursor                          # get-next-record.go:18 -> github.com/joel-ling/lmdb-bdd/test.getNextRecord
    Then I should see an error "MDB_NOTFOUND"                            # see-error.go:19 -> github.com/joel-ling/lmdb-bdd/test.seeError
```
```txt
7 scenarios (7 passed)
66 steps (66 passed)
26.43273ms
PASS
ok  	github.com/joel-ling/lmdb-bdd	0.032s
```
