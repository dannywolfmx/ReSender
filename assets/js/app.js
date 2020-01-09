document.addEventListener('DOMContentLoaded', (event) => {
		
	//Boton de buscar
	const botonAgregarCliente = document.getElementById("agregarCliente");
	
	botonAgregarCliente.addEventListener("click", ()=>{
		let form = document.getElementById("formAgregarCliente")
		
		form = new FormData(form);
		fetch('./clientes',{
			method:'POST',
			body:form
		})

	})

})


