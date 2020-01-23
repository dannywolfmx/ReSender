class FormCliente extends HTMLFormElement{
	constructor(){
		super()
		this._botonEnviar = document.createElement("input")
		this._botonEnviar.type = "submit"
		this._botonEnviar.className = "stack icon-paper-plane"
	}

	_render(){
			
		this.innerHTML = `
			Agregar Cliente:
			<input type="number" name="id" placeholder="Agregar id"/>
			<input type="text" name="nombre" placeholder="Agregar cliente"/>
		`
	}

	//Enviar el formulario al servidor
	enviarFormulario(){
		fetch('./clientes',{
			method:'POST',
			body:new FormData(this)
		}).then((respuesta) => {
			//Limpiar el formulario si la respuesta es Ok
			if(respuesta.ok){
				this.reset()
			}else{
				//Notificar error en el formato
				alert("Error: revisa tus datos")
			}
		}).catch((error) =>{
			//Posible error en la conexion
			console.log("Error de conexion")
		})
	}

	connectedCallback(){
		//Accion al enviar formulario
		 this._render()	
		 this.appendChild(this._botonEnviar)
		 this._botonEnviar.addEventListener("click", (e)=>{
			e.preventDefault()
			this.enviarFormulario()	
		})
	}
}

customElements.define("form-cliente", FormCliente, {extends: "form"})
