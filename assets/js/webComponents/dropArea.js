class DropArea extends HTMLElement{
	constructor(){
		super()
		this.attachShadow({mode:'open'})
		this. notificacion = ''
	}

	connectedCallback(){
		this._render()		
		let notificacion = this.shadowRoot.getElementById("notificacion")	
		//Evento que indica que el elemento fue arrastrado al formulario
		this.addEventListener('drop',(e)=>{
			e.preventDefault()
			let items = e.dataTransfer.items
			if(items){
				let archivos = this._recuperaArchivos(items)
				//Enviar los archivos como un evento
				this.dispatchEvent(new CustomEvent('archivos',{
					"detail":{archivos}
				}))
			}else{
				console.log("No se encontraron archivos")
			}
			notificacion.style.display = "none"

		})

		//Evento para mostrar que el elemento esta siendo arrastrado
		//al formulario
		this.addEventListener('dragover',(e)=>{
			e.preventDefault()
			//Prograr animacion de drag y deteccion de archivos validos
			notificacion.style.display = "block"
		})
	}

	_render(){
		let style = `
			<style>
				#dropArea{
					height:100px;
					width: 100%;
					margin:10px;
					border:1px solid black;
				}

				#notificacion{
					display:none;	
				}

				.file{
					border: 1px solid black;
					display: inline-block;
					padding: 10px;
					margin: 10px;
				}

			</style>
		`
		this.shadowRoot.innerHTML = `
			${style}
			<div id="dropArea">
				<div id="notificacion">
					<h1>Suelta los archivos</h1>
				</div>
			</div>
		`
	}
	
	//Recuperar archivos de un dataTransfer
	_recuperaArchivos(items){
		let archivos = []
		
		for(let i=0; i< items.length; i++){
			//Detectar si de tipo file
			if(items[i].kind === 'file'){
				let archivo = items[i].getAsFile();
				archivos.push(archivo)
				this._crearArchivo(archivo.name)
			}
		}
		return archivos
	}
	
	_crearArchivo(nombre){
		let contenedor = document.createElement("div")
		contenedor.className = "file"
		contenedor.innerText = nombre
		let dropArea = this.shadowRoot.getElementById('dropArea')
		dropArea.appendChild(contenedor)
	}
}

customElements.define("drop-area", DropArea )
