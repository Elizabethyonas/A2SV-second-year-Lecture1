import pytest
from main import fizzbuzz

def test_fizzbuzz():
    assert fizzbuzz(1) == "1"
    assert fizzbuzz(3) == "Fizz"
    assert fizzbuzz(5) == "Buzz"
    assert fizzbuzz(15) == "FizzBuzz"
    assert fizzbuzz(30) == "FizzBuzz"
    assert fizzbuzz(7) == "7"
    assert fizzbuzz(9) == "Fizz"
    assert fizzbuzz(10) == "Buzz"
    assert fizzbuzz(45) == "FizzBuzz"

if __name__ == "__main__":
    pytest.main()
