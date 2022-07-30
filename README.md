# Owner's Note
The purpose of this project is for showing my code style and how I design a project stucture including frontend using Vue and backend using Golang via Eshop Mock.
The Eshop Mock mimics E-commerce web application but I have just created 1 service. It is a catalog service which show products from database (Postgrest). In the future
I will add another services. I think I will apply microservice but now it has just 1 microservice.

- FrontEnd (VUE) - I try to apply advantage of component which easy to reuse so in my project I write code which frequently reuse to component. I implement Vuex to mange
my data and keep following [Vue Style](https://v2.vuejs.org/v2/style-guide/?redirect=true) Guide.

## Project Structure
- Backend (Golang) - I apply Hexagonal architecture for my project because it is easy to change package in the future. The concept of Hexagonal is intering. It manges 
package to interface which connect to other interface so it pretty simple to change interface when you want to update your package. In my project I implement GIN and GORM
to interface. It not only easy for changing but also simple for writing unit test. More over I write test cases for my CURD API. I use [gomock](https://github.com/golang/mock)
to generate test in interface for me.

## Built With
- [VUE](https://vuejs.org/) -The frontend framwork
- [GIN](https://github.com/gin-gonic/gin) - The web framework
- [GORM](https://gorm.io/) - ORM library

## Reference (My conduct code)
 - [go-gin-api](https://github.com/xinliangnote/go-gin-api)-Xinliangnote
 - [Vue-The Complete Guide](https://www.udemy.com/course/vuejs-2-the-complete-guide)-Maximilian Schwarzmüller
 - [Effective Go: Architecture and Design Patterns](https://skooldio.com/courses/effective-go)-Pallat Anchaleechamaikorn
