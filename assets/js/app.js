document.addEventListener('DOMContentLoaded', (event) => {
		
	//Boton de buscar
	const botonAgregarCliente = document.getElementById("agregarCliente");
	
	botonAgregarCliente.addEventListener("click", (e)=>{
		e.preventDefault()

		let form = document.getElementById("formAgregarCliente")
		
		fetch('./clientes',{
			method:'POST',
			body:new FormData(form)
		}).then((respuesta) => {
			//Limpiar el formulario si la respuesta es Ok
			if(respuesta.ok){
				form.reset()
			}else{
				//Notificar error en el formato
				alert("Error: revisa tus datos")
			}
		}).catch((error) =>{
			//Posible error en la conexion
			console.log("Error de conexion")
		})

	})

})


