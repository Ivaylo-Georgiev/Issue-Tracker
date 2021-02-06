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

// Client is used to make transactions in the database
var Client mongo.Client

const (
	connectionURI      = "mongodb://localhost/issuetracker"
	dbName             = "issuetracker"
	usersCollection    = "users"
	projectsCollection = "projects"
	issuesCollection   = "issues"
)

// Connect establishes a connection to the database
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
}

// IsertRegisteredUser inserts a new user in the 'users' collection
func InsertRegisteredUser(registeredUser user.User) {
	collection := Client.Database(dbName).Collection(usersCollection)
	_, err := collection.InsertOne(context.TODO(), registeredUser)
	if err != nil {
		log.Fatal(err)
	}
}

// FindRegisteredUser checks whether a specific username is already in the 'users' collection
func FindRegisteredUser(username string) (user.User, error) {
	collection := Client.Database(dbName).Collection(usersCollection)
	filter := bson.M{"username": username}
	var registeredUser user.User
	err := collection.FindOne(context.TODO(), filter).Decode(&registeredUser)

	return registeredUser, err
}

// InsertNewProject inserts a new project in the 'projects' collection
func InsertNewProject(newProject project.Project) {
	collection := Client.Database(dbName).Collection(projectsCollection)
	_, err := collection.InsertOne(context.TODO(), newProject)
	if err != nil {
		log.Fatal(err)
	}
}

// FindExistingProject checks whether a specific project is already in the 'projects' collection
func FindExistingProject(name string) (project.Project, error) {
	collection := Client.Database(dbName).Collection(projectsCollection)
	filter := bson.M{"name": name}
	var existingProject project.Project
	err := collection.FindOne(context.TODO(), filter).Decode(&existingProject)

	return existingProject, err
}

// InsertNewIssue inserts a new issue in the 'issues' collection
func InsertNewIssue(newIssue issue.Issue) {
	collection := Client.Database(dbName).Collection(issuesCollection)
	_, err := collection.InsertOne(context.TODO(), newIssue)
	if err != nil {
		log.Fatal(err)
	}
}

// FindExistingIssue checks whether an issue with a title and a project is already in the 'issues' collection
func FindExistingIssue(project string, title string) (issue.Issue, error) {
	collection := Client.Database(dbName).Collection(issuesCollection)
	filter := bson.M{"project": project, "title": title}
	var existingIssue issue.Issue
	err := collection.FindOne(context.TODO(), filter).Decode(&existingIssue)

	return existingIssue, err
}

// ResolveIssue updates an entry in the 'issues' collection by changing the value of the 'resolved' attribute to 'true'
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

// ListIssues creates a comma separated list of the names of all issues in a project
func ListIssues(project string) []issue.Issue {
	collection := Client.Database(dbName).Collection(issuesCollection)
	cursor, _ := collection.Find(
		context.TODO(),
		bson.M{"project": project})

	var issues []issue.Issue
	cursor.All(context.TODO(), &issues)

	return issues
}
