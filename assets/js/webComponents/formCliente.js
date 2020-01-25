class FormCliente extends HTMLFormElement{
	constructor(){
		super()
		this._botonEnviar = document.createElement("button")
		this._botonEnviar.type = "submit"
		this._botonEnviar.className = "btn btn-primary"
		this._botonEnviar.innerText = "Guardar"
	}

	_render(){
			
		this.innerHTML = `
			<div class="form-group">
				<label for="id">
					Id (Elemento temporal)
				</label>
				<input type="number" name="id" placeholder="Agregar id" class="form-control"/>
			</div>
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
		
	//Recuperar archivos de un dataTransfer
	_recuperaArchivos(items){
		let archivos = []
		for(let i=0; i< items.length; i++){
			//Detectar si de tipo file
			if(items[i].kind === 'file'){
				let archivo = items[i].getAsFile();
				archivos.push(archivo)
			}
		}
		return archivos
	}

	connectedCallback(){
		//Accion al enviar formulario
		 this._render()	
		 this.appendChild(this._botonEnviar)
		 this._botonEnviar.addEventListener("click", (e)=>{
			e.preventDefault()
			this.enviarFormulario()	
		})
		
		//Evento que indica que el elemento fue arrastrado al formulario
		this.addEventListener('drop',(e)=>{
			e.preventDefault()
			let items = e.dataTransfer.items
			if(items){
				let archivos = this._recuperaArchivos(items)
			}else{
				console.log("No se encontraron archivos")
			}
			console.log('drop')
		})

		//Evento para mostrar que el elemento esta siendo arrastrado
		//al formulario
		this.addEventListener('dragover',(e)=>{
			e.preventDefault()
			//Prograr animacion de drag y deteccion de archivos validos
		})
	}
}

customElements.define("form-cliente", FormCliente, {extends: "form"})
