import { useState } from 'react'
import './App.css'

function App() {
  const openArch =(e)=>{
    alert('Deseas abrir un archivo?')
  }
  return (
    <>
     <nav className="navBar">
      <img src="src/assets/frankestein.png" alt="" />
      <img src="/src/assets/archivo.png" className="imgArch" alt="abrir archivo" onClick={openArch} />
     </nav>
      <div className='consoles'>
        <textarea className="txt_entrada" name="txt_entrada" id=""></textarea>
        <textarea className='txt_salida'  name="txt_salida" id="" readOnly></textarea>
      </div>
      <div className="card">
        <button className='btnA'>EJECUTAR</button>
      </div>
    </>
  )
}

export default App
