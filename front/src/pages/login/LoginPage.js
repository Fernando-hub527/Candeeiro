import "../../static/css/styleLoginPage.css"
import "@fontsource/comfortaa"
import arrowIcon from "../../static/img/arrowIcon.svg"
import eyeIcon from "../../static/img/eyeIcon.svg"
import linkedinIcon from "../../static/img/linkedinIcon.png"
import googleIcon from "../../static/img/googleIcon.png"

import { useState } from "react"
import { useNavigate } from "react-router-dom"
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

export function LoginPage(props) {
    const [passwordVisibility, setVisibility] = useState(false)
    const navigate = useNavigate();

    const login = (e) => {
        toast("Usuário ou senha inválido", {type: "error"});
        e.preventDefault()

    }
    

    return <div className="page_login">
        <div className="login_header">
            <img onClick = {() => navigate("/welcome")} src={arrowIcon} alt="Um icone preto de uma seta apontada para a esquerda"/>
            <h2>Login</h2>
        </div>
        <form>
            <div className="form_input form_user">
                <input type="text" id="nickName" placeholder="Usuário" />
            </div>
            <div className="form_input form_password">
                <input type={passwordVisibility ? "text":"password"} id="password" placeholder="Senha" />
                <div className="form_password_eye" onClick={()=>setVisibility((state)=> !state)}>
                    <hr className={`form_password_eye-${passwordVisibility ? "visible" : "invisible"}`}/>
                    <img draggable={false} src={eyeIcon} alt="Icone roxo de um olho indicando se a senha está visível"/>
                </div>
            </div>
            <button onClick={login} className="btn btn_input">Entrar</button>
        </form>
        <div className="form_division">
            <hr />
            <p>ou</p>
            <hr />
        </div>

        <div className="login_options">
            <button className="btn login_option_google"><img src={googleIcon}/><p>Continuar com google</p></button>
            <button className="btn login_option_linkedin"><img src={linkedinIcon}/><p>Continuar com linkedin</p></button>

        </div>
        <ToastContainer limit={1}/>
    </div>


}