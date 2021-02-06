package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.fmi/issuetracker/issue"
	"go.fmi/issuetracker/project"
	"go.fmi/issuetracker/user"
)

var Client mongo.Client

const (
	connectionURI      = "mongodb://localhost/issuetracker"
	dbName             = "issuetracker"
	usersCollection    = "users"
	projectsCollection = "projects"
	issuesCollection   = "issues"
)

func Connect() {
	clientLocal, err := mongo.NewClient(options.Client().ApplyURI(connectionURI))
	Client = *clientLocal
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = clientLocal.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//defer Client.Disconnect(ctx)
}

func InsertRegisteredUser(registeredUser user.User) {
	collection := Client.Database(dbName).Collection(usersCollection)
	_, err := collection.InsertOne(context.TODO(), registeredUser)
	if err != nil {
		log.Fatal(err)
	}
}

func FindRegisteredUser(username string) (user.User, error) {
	collection := Client.Database(dbName).Collection(usersCollection)
	filter := bson.M{"username": username}
	var registeredUser user.User
	err := collection.FindOne(context.TODO(), filter).Decode(&registeredUser)

	return registeredUser, err
}

func InsertNewProject(newProject project.Project) {
	collection := Client.Database(dbName).Collection(projectsCollection)
	_, err := collection.InsertOne(context.TODO(), newProject)
	if err != nil {
		log.Fatal(err)
	}
}

func FindExistingProject(name string) (project.Project, error) {
	collection := Client.Database(dbName).Collection(projectsCollection)
	filter := bson.M{"name": name}
	var existingProject project.Project
	err := collection.FindOne(context.TODO(), filter).Decode(&existingProject)

	return existingProject, err
}

func InsertNewIssue(newIssue issue.Issue) {
	collection := Client.Database(dbName).Collection(issuesCollection)
	_, err := collection.InsertOne(context.TODO(), newIssue)
	if err != nil {
		log.Fatal(err)
	}
}

func FindExistingIssue(project string, title string) (issue.Issue, error) {
	collection := Client.Database(dbName).Collection(issuesCollection)
	filter := bson.M{"project": project, "title": title}
	var existingIssue issue.Issue
	err := collection.FindOne(context.TODO(), filter).Decode(&existingIssue)

	return existingIssue, err
}

func ResolveIssue(project string, title string) {
	collection := Client.Database(dbName).Collection(issuesCollection)
	_, err := collection.UpdateOne(
		context.TODO(),
		bson.M{"project": project, "title": title},
		bson.D{
			{"$set", bson.D{{"resolved", "true"}}},
		},
	)

	if err != nil {
		log.Fatal(err)
	}
}
