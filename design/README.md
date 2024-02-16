# System Design Exercise: High-Level Redesign of Plagiarism Detector for CodeGrade

## Introduction

This exercise involves creating a high-level design for the redesign
of CodeGrade's existing plagiarism detection system, which currently
employs JPlag. The primary aim is to enhance the overall system's
performance and introduce real-time feedback for users during the
detection process. It's important to maintain the integrity of the
plagiarism detection results, keeping them identical to the current
implementation.

JPlag works on first parsing submissions to a token stream, using the
abstract `Parse(Submission): ParsedSubmission` function, and then
comparing two submissions by `Compare(ParsedSubmission,
ParsedSubmission): Similarity`. These two functions are pure, and both
take considerable amount of time, with `Parse` being the slower of the
two. Two find plagiarism in a assignment every submission needs to be
compared to every other submission.

## Scope of the Exercise

- **Focus**: High-level architectural design. **No coding or implementation is needed**.
- **Time Allocation**: Spend no more than 4 hours on this task.
- **Key Goals**: Address system architecture, user feedback, performance (excluding individual function optimization), scalability, and reliability.

## Existing System Overview

- **Monolithic Design**: The current system is monolithic, leading to performance issues and lack of real-time feedback.
- **Function Optimization**: Assume that the individual functions (`Compare` and `Parse`) are already optimized and cannot be altered for further performance gains.
- **Result Consistency**: The redesigned system must produce results that are identical to the current JPlag implementation.

## Core Aspects of the Redesign

### 1. **System Architecture**

Propose a new architectural design that can enhance the system's efficiency and scalability.

### 2. **Real-Time Feedback Mechanism**

Design a system feature that provides educators with real-time updates on the status of plagiarism checks.

### 3. **Optimizing System Performance**

- Focus on system-level optimizations that can improve overall performance, such as load balancing, parallel processing of submissions, and efficient resource management.
- Suggest strategies for handling large volumes of data and requests, ensuring timely processing.

### 4. **Scalability and Reliability**

- The system should be capable of handling varying loads, especially during peak submission times.
- Include considerations for fault tolerance, error handling, and ensuring system reliability.

## Deliverables

**Design Documentation**: Your submission should consist of a design document. This document should outline your proposed system architecture and explain how it addresses the key goals of enhancing performance and providing real-time feedback. It should also illustrate how your design integrates with the existing `Compare` and `Parse` functions while maintaining the accuracy and consistency of plagiarism detection. The format and structure of the document are at your discretion, allowing you to present your design in a way that best communicates your ideas.

## Submission Instructions

- Submit a document detailing your high-level system design.
- Clearly state any assumptions made during the design process.
- Your design should maintain the accuracy and consistency of the existing plagiarism detection results.
- You may assume any infrastructure offered by any public cloud offering.
