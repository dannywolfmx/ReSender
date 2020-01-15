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

	//Listar clientes en la vista
	const clientesDIV = document.getElementById("clientes")
	
	fetch('clientes').then((r)=>{
		return r.json()
	}).then((clientes) =>{
		//Agregar clientes a la vista html
		clientes.map(c => {
			let div = document.createElement("div")
			let span = document.createElement("span")
			let text = document.createTextNode(c.nombre)
			span.appendChild(text)
			div.appendChild(span)
			clientesDIV.appendChild(div)
		})
	}).catch((error)=>{
		console.log("Error al obtener clientes: ", error )
	})

})


