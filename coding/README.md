# Coding Exercise

## Project Overview

This take-home project is designed as a simple service allowing users to post messages on a board, similar to a course's homepage. Each post has an expiry time, set at the time of creation, and posts are automatically hidden from the board after this time elapses. The goal of this project is to demonstrate basic full-stack development skills.

### Time Allocation

Candidates are expected to spend no more than 5 hours on this task. It's acceptable to submit an incomplete solution, but the submitted code should be runnable, and any gaps in implementation should be acknowledged and documented.

## Project Structure

The project is divided into two main parts:

- **Backend**: Located in the `backend` folder. It's built using Go, Go Fiber, and Ent.

  - [Go (Golang)](https://golang.org/)
  - [Go Fiber](https://gofiber.io/)
  - [Ent](https://entgo.io/)

- **Frontend**: Located in the `frontend` folder. It's developed using Vue 3, TSX, and Pinia.
  - [Vue 3](https://v3.vuejs.org/)
  - [TSX (TypeScript JSX)](https://www.typescriptlang.org/docs/handbook/jsx.html)
  - [Pinia](https://pinia.vuejs.org/)

## Setup and Running the Application

### Backend

Navigate to the `backend` directory. You can run it using the provided `Makefile` in this directory.

### Frontend

1. **Install Dependencies**: Ensure that you have `pnpm` installed. If not, install it via npm:

   ```bash
   npm install -g pnpm
   ```

2. Navigate to the `frontend` directory:

   ```bash
   cd frontend
   ```

3. Install the required packages using `pnpm`:

   ```bash
   pnpm install
   ```

4. Run the frontend application:

   ```bash
   pnpm run serve
   ```

   This will start the Vue 3 application, typically available at `http://localhost:8080`.

## Project Requirements and Expectations

### Flexibility in Implementation

While this project comes with a specific set of tools and
technologies, you are free to make changes to the technology stack,
architecture, or any other aspect of the project. We encourage
innovation and creativity, so feel free to take the project in a
direction you find interesting or more effective. However, with each
change or deviation from the provided setup, please ensure to:

- **Provide a Clear Justification**: Document your reasons for making the change. Explain why your approach is more effective, efficient, or better suited to the project goals.
- **Ensure Functionality**: Regardless of the changes, the basic functionality of posting and viewing messages with expiry times should be maintained.
- **Maintain Code Quality**: Your code should be clean, well-organized, and follow best practices in software development. Good code structure and readability are important.
- **Document Your Work**: Clearly document any deviations from the original specifications, including changes in technology, architecture, or design. Provide sufficient comments and README updates to help understand your approach.

## Submission

Please submit your code by either providing a GitHub link, or by
emailing thomas@codegrade.com.

Your submission should reflect both your technical skills and your
ability to think critically about software design choices. We are as
interested in your reasoning and decision-making process as we are in
your coding skills. Ensure that your final submission is a runnable
version of the application, with any gaps or incomplete features
clearly documented.
