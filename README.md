# ReSender
Sistema de tracking y envio de notificaciones de pedidos

Correr aplicación

go run main.go

Esta aplicacion pretende seguir los lineamientos de la arquitectura hexagonal.

## Descripcion de las carpetas

* app
> App representa el codigo fuente principal de la aplicación
 * delivery
 > Delivery la capa exterior de nuestra arquitectura, es la parte que se comunica con el mundo exterior, en este caso contiene la api json
 * domain
 > Domain se encuentra nuestra logica de negocio, en esta esta los modelos y servicios de la aplicacion
 > > Este es el nucleo de nustra aplicación.
* assets
* auth 
> Auth es un middleware, esta no es dependiende de la aplicación, por lo que puede ser reutilizado en otros proyectos 
