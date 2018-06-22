package main

import "testing"

func TestCheckForValidSite(t *testing.T) {
  validItem := Item{"https://google.com", false}
  validItem.check()
  if validItem.Status != true {
    t.Errorf("Expected true Status for a valid site ")
  }
}

func TestCheckForInvalidSite(t *testing.T) {
  validItem := Item{"https://google123456789.com", false}
  validItem.check()
  if validItem.Status != false {
    t.Errorf("Expected false Status for a invalid site ")
  }
}
