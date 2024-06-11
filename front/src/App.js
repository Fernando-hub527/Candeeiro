import { Route, Routes } from 'react-router-dom';
import {PrivateRoute} from "./services/auth/PrivateRout"
import {LoginPage} from "./pages/login/LoginPage"
import {HomePage} from "./pages/home/HomePage"
import { Presentation } from './pages/presentation/Presentation';
import { Register } from './pages/register/Register';

function App() {
  return(
      <Routes>
        <Route path="*" element={<Presentation />}/>
        <Route path="/login" element={<LoginPage />}/>
        <Route path="/register" element={<Register />}/>
        <Route path="/home" element={<PrivateRoute><HomePage /></PrivateRoute>}/>
      </Routes>
  )
}

export default App;