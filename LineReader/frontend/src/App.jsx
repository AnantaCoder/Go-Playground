import { useState } from '../$node_modules/@types/react/index.js'
import './App.css'
import axios from '../$node_modules/axios/index.js'

function App() {
  const{file,
    result,
    loading,
    error,
    handleFileChange,
    handleSubmit,
    cancelUpload}= customReactQuery("/upload")

  return (
    <>
        <h1>File Handler.com</h1>
        <h2>Upload your files here 📂</h2>
        <input type="file" onChange={handleFileChange} />
        <button onClick={handleSubmit} disabled={!file || loading}>Send</button>
        {loading && <button onClick={cancelUpload}>Cancel</button>}
        
        {error && <div style={{ color: 'red' }}>{error}</div>}
        {result && (
          <div >
            <p>📄Lines: {result.lines}</p>
            <p>🔤Words:{result.words} </p>
            <p>🔡Characters:{result.characters} </p>
             </div>
        )}
    </>
  )
}

const customReactQuery = (endpoint)=>{

  const [ file, setFile] = useState(null)
  const [result,setResult] = useState(null)
  const [loading,setLoading] = useState(false)
  const [ error,setError]  = useState(null)
  const [controller, setController] = useState(null);

  const SERVER_URL = "http://localhost:8080"
  const urlPath = `${SERVER_URL}${endpoint}`


  const handleFileChange = (e) => {
    setResult(null)
    setError(null)
    setFile(e.target.files[0])
  }
  const cancelUpload= ()=>{
    if(controller){
      controller.abort()
      setLoading(false)
    }
  }
  const handleSubmit = async()=>{
    if(!file){
      setError("File is not Uploaded")
      return 
    }
    // cancel any prev ongoing request
    if (controller){
      controller.abort();
    }

    const newController = new AbortController()
    const signal = newController.signal;

    try {
      setLoading(true)
      setError(null)

      //create from data for file upload 
      const formData = new FormData()
      formData.append('file',file)

      const response = await axios.post(
        urlPath, formData,{headers:{"Content-Type": "multipart/form-data"},signal}
      )
      setResult(response.data)
      setLoading(false)


    } catch (error) {
      if(axios.isCancel(error)){
        setError("Request Was Cancelled. ")
      }else{
        console.error("Error details:", error);
        setError(error.response?.data?.message || "Error Connecting to the server.")
      }
      setLoading(false)
    }

  }
  return{
    file,result,loading,error,cancelUpload,handleFileChange,handleSubmit,setController,setError,setFile,setLoading,
  }
}

export default App
