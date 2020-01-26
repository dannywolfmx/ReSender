
class FormOrden extends HTMLFormElement{
	constructor(){
		super()
		this._botonEnviar = document.createElement("button")
		this._botonEnviar.type = "submit"
		this._botonEnviar.className = "btn btn-primary"
		this._botonEnviar.innerText = "Guardar"

		this.error = ""
	}

	_render(){
		let formulario = 			
		`
			<div class="form-group">
				<label for="Factura">
					Factura
				</label>
				<input type="text" name="factura" placeholder="Numero de factura" class="form-control"/>
			</div>
			<div class="form-group">
				<label for="Orden">
					Orden de compra
				</label>
				<input type="text" name="orden" placeholder="Orden de compra" class="form-control"/>
			</div>
			<div is="drop-area"></div>
			<drop-area id="dropArea"></drop-area>
		`

		this.innerHTML = `
			${formulario}
		` 
	}
	

	//Enviar el formulario al servidor
	//Pendiente agregar funcionalidad para este form
	enviarFormulario(){
	//	fetch('./clientes',{
	//		method:'POST',
	//		body:new FormData(this)
	//	}).then((respuesta) => {
	//		//Limpiar el formulario si la respuesta es Ok
	//		if(respuesta.ok){
	//			this.reset()
	//			//Mencionar que se a creado un nuevo elemento
	//			this.dispatchEvent(new CustomEvent('crear'))
	//		}else{
	//			//Notificar error en el formato
	//			this.dispatchEvent(new CustomEvent('error'))
	//			alert("Error: revisa tus datos")
	//		}
	//	}).catch((error) =>{
	//		//Posible error en la conexion
	//		this.dispatchEvent(new CustomEvent('sin-conexion'))
	//		console.log("Error de conexion")
	//	})
	}
		

	connectedCallback(){
		//Accion al enviar formulario
		 this._render()	
		 this.appendChild(this._botonEnviar)
		 this._botonEnviar.addEventListener("click", (e)=>{
			e.preventDefault()
			this.enviarFormulario()	
		})
		
		document.getElementById("dropArea").addEventListener("archivos",(e)=>{
			console.log(e.detail.archivos)
		})
		
	}
}

customElements.define("form-orden", FormOrden, {extends: "form"})
