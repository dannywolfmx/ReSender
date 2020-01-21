class FormCliente extends HTMLFormElement{
	constructor(){
		super()
		this.template()	
	}

	template(){
		this.innerHTML = `
				Agregar Cliente:
					
				<input type="number" name="id" placeholder="Agregar id"/>
				<input type="text" name="nombre" placeholder="Agregar cliente"/>
				<input type="submit" id="agregarCliente" class="stack icon-paper-plane"/>
		`
	}

	connectedCallback(){
		 //Accion al enviar formulario

		 this.querySelector("#agregarCliente").addEventListener("click", (e)=>{
			e.preventDefault()
			this.submit()
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

customElements.define("form-cliente", FormCliente, {extends:"form"})
