Feature: LMDB
  In order to verify the effects of software interacting with LMDB
  As a software engineer practising behaviour-driven development
  I need to establish some baseline observations

  Scenario: Get non-existent record
    Given there is a new LMDB environment "env"
    When I get a record "key" in DB "deebee" of environment "env"
    Then I should see an error "MDB_NOTFOUND"
    And the ID of last committed transaction in environment "env" should be 0

  Scenario: Put record
    Given there is a new LMDB environment "env"
    When I put a record "key" "value" in DB "deebee" of environment "env"
    Then the ID of last committed transaction in environment "env" should be 1

  Scenario: Put and then get record
    Given there is a new LMDB environment "env"
    When I put a record "key" "value" in DB "deebee" of environment "env"
    And I get a record "key" in DB "deebee" of environment "env"
    Then I should see a value "value"
    And I should see no error
    And the ID of last committed transaction in environment "env" should be 1

  Scenario: Get record put in different environment
    Given there is a new LMDB environment "primary"
    And there is a new LMDB environment "secondary"
    And there is a record "key" "value" in DB "DB" of environment "primary"
    When I get a record "key" in DB "DB" of environment "secondary"
    Then I should see an error "MDB_NOTFOUND"

  Scenario: Get records using a cursor
    Given there is a new LMDB environment "env"
    And there is a record "41" "A" in DB "ASCII" of environment "env"
    And there is a record "43" "C" in DB "ASCII" of environment "env"
    And there is a record "42" "B" in DB "ASCII" of environment "env"
    When I open a cursor to DB "ASCII" of environment "env"
    And I get the next record using the cursor
    Then I should see a record "41" "A"
    And I should see no error
    When I get the next record using the cursor
    Then I should see a record "42" "B"
    And I should see no error
    When I get the next record using the cursor
    Then I should see a record "43" "C"
    And I should see no error
    When I get the next record using the cursor
    Then I should see an error "MDB_NOTFOUND"

  Scenario: Get records using a cursor, resetting and renewing read transaction
    Given there is a new LMDB environment "env"
    And there is a record "41" "A" in DB "ASCII" of environment "env"
    And there is a record "42" "B" in DB "ASCII" of environment "env"
    And there is a record "43" "C" in DB "ASCII" of environment "env"
    When I open a cursor to DB "ASCII" of environment "env"
    And I get the next record using the cursor
    Then I should see a record "41" "A"
    And I should see no error
    When I get the next record using the cursor
    Then I should see a record "42" "B"
    And I should see no error
    When I reset the transaction relating to the cursor
    And I renew the transaction relating to the cursor, and the cursor
    And I get the next record using the cursor, specifying the previous key
    Then I should see a record "43" "C"
    And I should see no error

  Scenario: List databases in environment
    Given there is a new LMDB environment "env"
    And there is a record "key" "value" in DB "db0" of environment "env"
    And there is a record "key" "value" in DB "db1" of environment "env"
    And there is a record "key" "value" in DB "db2" of environment "env"
    When I open a cursor to the root database of environment "env"
    And I get the next record using the cursor
    Then I should see a key "db0"
    And I should see no error
    When I get the next record using the cursor
    Then I should see a key "db1"
    And I should see no error
    When I get the next record using the cursor
    Then I should see a key "db2"
    And I should see no error
    When I get the next record using the cursor
    Then I should see an error "MDB_NOTFOUND"
