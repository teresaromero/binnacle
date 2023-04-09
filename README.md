# Binnacle

Binnacle is an app where you can store all your travel notes and share them with anyone you like.

## Motivation and project progress

During my last vacation i wanted to share with a determined number of people insights, pictures and stories regarding my traveling. 

I used Instagram for this matter with a private account so i can control the scope of people being able to access my posts and updates. 

Problem was, some folks don't have Instagram. Also, after the vacation, some people also ask about details on budget, planning and places to visit.

From a product perspective, the purpose of this project is to MVP a simple app to track travel plans and being able to share them privately without the requirement of having an account.

From a dev perspective, the purpose is to have a side-project where to find, experiment and learn.

## Tech

### Server

Binnacle server is aim to be a gRPC server serving multiple services, as i wanted to explore gRPC implementation.

First "challenge" to solve would be to have a common `proto` repository where to store the definition of the services, so in the "near" future i can share them between server an client. For the time being this "repository" will be a folder on the repo and some `make` commands to control the protobuf generation.

The gRPC server is enabled with `reflection` as i will be using Postman to test the API and i want the services to load from the source, instead of importing a file. 

## License

[MIT](https://choosealicense.com/licenses/mit/)

---

Coded with ❤️ - Copyright © 2023 [Teresa Romero](https://github.com/teresaromero)