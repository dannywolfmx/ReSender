class FormCliente extends HTMLElement{
	constructor(){
		super()
		this.attachShadow({mode:"open"})
		this.template()	
	}

	template(){
		this.shadowRoot.innerHTML = `
			<link rel="stylesheet" href="https://unpkg.com/picnic">
			<form>
				Agregar Cliente:
				<input type="number" name="id" placeholder="Agregar id"/>
				<input type="text" name="nombre" placeholder="Agregar cliente"/>
				<input type="submit" id="boton" class="stack icon-paper-plane"/>
			</form>
		`
	}

	connectedCallback(){
		//Accion al enviar formulario
		 this.shadowRoot.getElementById("boton").addEventListener("click", (e)=>{
			e.preventDefault()
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
			
		})
	}
}

customElements.define("form-cliente", FormCliente)
