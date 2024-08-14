# Quiz App

## Description
Quiz App is a simple CLI (Command-Line Interface) application that allows users to take quizzes directly from their terminal.

## Usage
To install and run the Quiz App, follow these steps:

1. Before playing the quiz, you need to start the API. Run the following command:

   ```bash
   go run main.go api

1. Once the API is running, you can start the quiz by running:

   ```bash
   go run main.go quiz


## Description

This solution features an open API with static token, which is an intentional choice to simplify the implementation.

The `oapi-codegen` library was used to generate the API handlers.

Mocks for testing were created using the `mockery` tool.

In a few places, tests are deliberately missing; I have written some tests to demonstrate my approach.


