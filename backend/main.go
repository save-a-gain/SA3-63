package main
 
import (
   "context"
   "log"
 
   "github.com/save-a-gain/app/controllers"
   _ "github.com/save-a-gain/app/docs"
   "github.com/save-a-gain/app/ent"
   "github.com/gin-contrib/cors"
   "github.com/gin-gonic/gin"
   _ "github.com/mattn/go-sqlite3"
   swaggerFiles "github.com/swaggo/files"
   ginSwagger "github.com/swaggo/gin-swagger"
)
// @title SUT SA Example API
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/
 
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
 
// @host localhost:8080
// @BasePath /api/v1
 
// @securityDefinitions.basic BasicAuth
 
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
 
// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information
 
// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information
 
// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information
 
// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

type Employees struct {
	Employee []Employee
}

type Employee struct {
	Employeename string
	Employeeemail string
	Password string
}

type Repairinvoices struct {
	Repairinvoice []Repairinvoice
}

type Repairinvoice struct {
	Symptomid int
	Deviceid int
	Userid int
	Statusrepairid int
}

type Statusts struct {
	Statust []Statust
}

type Statust struct {
	Statustname string
}

func main() {
   router := gin.Default()
   router.Use(cors.Default())
 
   client, err := ent.Open("sqlite3", "file:return.db?cache=shared&_fk=1")
   if err != nil {
       log.Fatalf("fail to open sqlite3: %v", err)
   }
   defer client.Close()
 
   if err := client.Schema.Create(context.Background()); err != nil {
       log.Fatalf("failed creating schema resources: %v", err)
   }
 
   v1 := router.Group("/api/v1")
   controllers.NewEmployeeController(v1, client)
   controllers.NewRepairinvoiceController(v1, client)
   controllers.NewStatustController(v1, client)
   controllers.NewReturninvoiceController(v1, client)
 
   // Set Employees Data
	employees := Employees{
		Employee: []Employee{
			Employee{"john","qwwe@gmail.com","1223"},
			Employee{"gain","sddf@gmail.com","2334"},
			Employee{"tod","zxxc@gmail.com","3445"},
		},
	}
	for _, em := range employees.Employee {
		client.Employee.
			Create().
			SetEmployeename(em.Employeename).
			SetEmployeeemail(em.Employeeemail).
			SetPassword(em.Password).
			Save(context.Background())
	}

	// Set Repairinvoices Data
	repairinvoices := Repairinvoices{
		Repairinvoice: []Repairinvoice{
			Repairinvoice{101, 201, 301, 1001},
			Repairinvoice{102, 202, 302, 1002},
			Repairinvoice{101, 203, 303, 1001},
		},
	}
	for _, rep := range repairinvoices.Repairinvoice {
		client.Repairinvoice.
			Create().
			SetSymptomid(rep.Symptomid).
			SetDeviceid(rep.Deviceid).
			SetUserid(rep.Userid).
			SetStatusrepairid(rep.Statusrepairid).
			Save(context.Background())
	}

	// Set Statusts Data
	statusts := Statusts{
		Statust: []Statust{
			Statust{"in delivery"},
			Statust{"completed"},
		},
	}
	for _, s := range statusts.Statust {
		client.Statust.
			Create().
			SetStatustname(s.Statustname).
			Save(context.Background())
	}


	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
 
