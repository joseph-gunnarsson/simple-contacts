## Overview

The simple-contact project is a command-line interface (CLI) tool for managing a contact list. It provides basic functionalities to add, update, delete, find, and list contacts. The contacts are stored locally and can be managed through various commands provided by this tool.

## Features

- **Add Contact:** Add a new contact with a name, phone number, and email address.
- **Update Contact:** Update an existing contact's name, phone number, or email address.
- **Delete Contact:** Remove a contact from the list by name.
- **Find Contact:** Retrieve and display details of a contact by name.
- **List Contacts:** Display all contacts stored in the system.

## Installation

To install the simple-contact CLI tool, first, clone the repository and build the tool:

```
git clone https://github.com/joseph-gunnarsson/simple-contacts.git
cd simple-contacts
go build -o simple-contact cmd/main.go
```

For easier access, you can move the compiled binary to your system's bin directory:

```
sudo mv simple-contact /usr/local/bin/
```

This allows you to run the tool from anywhere in your terminal by simply typing:

```
simple-contact
```

## Usage

The simple-contact tool supports several commands to manage your contacts. Below is a guide on how to use each command.

### 1. Add a Contact

To add a new contact:

```
simple-contact add --name "John Doe" --phone "123-456-7890" --email "johndoe@example.com"
```

### 2. Update a Contact

To update an existing contact's information:

```
simple-contact update --name "John Doe" --newName "Johnathan Doe" --phone "098-765-4321" --email "johnathan.doe@example.com"
```

> **Note:** You can update the name, phone number, and/or email address. You can provide one or more of these options.

### 3. Delete a Contact

To delete a contact by name:

```
simple-contact delete --name "John Doe"
```

### 4. Find a Contact

To find and display a contact by name:

```
simple-contact find --name "John Doe"
```

### 5. List All Contacts

To list all contacts stored in the system:

```
simple-contact list
```