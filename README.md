# ReSender
Sistema de tracking y envio de notificaciones de pedidos

Correr aplicación

go run main.go

Esta aplicacion pretende seguir los lineamientos de la arquitectura hexagonal.

## Descripcion de las carpetas

* app
   > App representa el codigo fuente principal de la aplicación
   * delivery - Delivery la capa exterior de nuestra arquitectura, es la parte que se comunica con el mundo exterior, en este caso contiene la api json
      * http/v1 - Verision 1 de la api hecha para http.
   * domain - Domain se encuentra nuestra logica de negocio, en esta esta los modelos y servicios de la aplicacion - Este es el nucleo de nustra aplicación.
   * repository - Implementación de la interface repository, esta implementación es la base de datos
      * gormrepo - repositorio para conectar por gorm (framework orm) a sqlite
      * mocks - un mocks es utilizado para realizar pruebas, no es usa una base de datos real, utiliza gomoks, para realizar pruebas mas facilmente.
      * sqlite - implementación de sqlite sobre el repositorio, sin usar un orm.
   * usecase - representa la implementación de la interface de caso de uso esta.
* assets
* auth - Auth es un middleware, esta no es dependiende de la aplicación, por lo que puede ser reutilizado en otros proyectos 
* registry/container.go - aqui se encuentran nuestro contenedor para la inyección de dependencias, utilizando el framework de "sarulabs/di"

prueba 3
