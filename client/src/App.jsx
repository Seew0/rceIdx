import React from 'react'
import CodeEditor from "./components/CodeEditor.jsx";
import { ToastContainer } from "react-toastify";
import Home from './Home.jsx';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
const App = () => {
  return (
    <>
      {console.log(process.env.REACT_APP_RAPID_API_URL)}
      <ToastContainer
        position="top-right"
        autoClose={2000}
        hideProgressBar={false}
        newestOnTop={false}
        closeOnClick
        rtl={false}
        pauseOnFocusLoss
        draggable
        pauseOnHover
      />

      <Router>
        <Routes>
          <Route exact path="/" element={<Home />} />
          <Route exact path="/codeeditor" element={<CodeEditor />} />
        </Routes>
      </Router>

    </>
  )
}

export default App