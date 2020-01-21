class FormCliente extends HTMLFormElement{
	constructor(){
		super()
		//this.attachShadow({mode:"open"})
		console.log("Hola")
		this.template()	
		//this.agregarEvento()
	}

	template(){
		this.innerHTML = `
				Agregar Cliente:
					
				<input type="number" name="id" placeholder="Agregar id"/>
				<input type="text" name="nombre" placeholder="Agregar cliente"/>
				<input type="submit" id="agregarCliente" class="stack icon-paper-plane"/>
		`
	}

	//Boton de buscar
	agregarEvento(){
		let botonAgregarCliente = this.shadowRoot.getElementById("agregarCliente"); 

		botonAgregarCliente.addEventListener("click", (e)=>{
			e.preventDefault()
	
			form = this.shadowRoot.getElementById("formulario")
			
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
	}
}

customElements.define("form-cliente", FormCliente, {extends:"form"})
