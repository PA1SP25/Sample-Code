/*
 * Author      : Luke Sathrum
 * Description : A Guessing Game in C++
 */

#include <iostream>
#include <cstdlib>
#include <ctime>
using std::cout;
using std::cin;
using std::endl;
using std::srand;
using std::rand;
using std::time;


// Program execution starts here
int main() {
  // Setup our variables
  int computers_guess, users_guess;
  // Setup the random number generator
  srand(time(0));
  // Get the random number.
  computers_guess = (rand() % 10) + 1;

  // Display a welcome message
  cout << "Welcome to the number guessing game.\n"
       << "I'm thinking of a number between 1 and 10.\n"
       << "What is your guess? ";
  // Get the user's guess
  cin >> users_guess;

  // Compare the numbers and give an appropriate message
  if (computers_guess == users_guess) {
    cout << "Good guess! You've won!\n";
  } else {
    cout << "I'm sorry that was not the correct number.\n"
         << "The number I was thinking of was " << computers_guess << ".\n";
  }

  // This ends the program
  return 0;
}
