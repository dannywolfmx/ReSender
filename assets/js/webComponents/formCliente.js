class FormCliente extends HTMLFormElement{
	constructor(){
		super()
		this._botonEnviar = document.createElement("button")
		this._botonEnviar.type = "submit"
		this._botonEnviar.className = "btn btn-primary"
		this._botonEnviar.innerText = "Guardar"

		this.error = ""
	}

	_render(){
		this.innerHTML = 			
		`
			<div class="form-group">
				<label for="Nombre">
					Nombre
				</label>
				<input type="text" name="nombre" placeholder="Agregar cliente" class="form-control"/>
			</div>
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
				//Mencionar que se a creado un nuevo elemento
				this.dispatchEvent(new CustomEvent('crear'))
			}else{
				//Notificar error en el formato
				this.dispatchEvent(new CustomEvent('error'))
				alert("Error: revisa tus datos")
			}
		}).catch((error) =>{
			//Posible error en la conexion
			this.dispatchEvent(new CustomEvent('sin-conexion'))
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