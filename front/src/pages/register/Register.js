import { useNavigate } from "react-router-dom"
import "../../static/css/styleLoginPage.css"
import arrowIcon from "../../static/img/arrowIcon.svg"
import eyeIcon from "../../static/img/eyeIcon.svg"
import linkedinIcon from "../../static/img/linkedinIcon.png"
import googleIcon from "../../static/img/googleIcon.png"
import { useState } from "react"
import { ToastContainer, toast } from "react-toastify"

export function Register(props){
    const [passwordVisibility, setVisibility] = useState(false)
    const navigate = useNavigate()
    const cadastrar = (e) => {
        toast("Seu usuário foi criado com sucesso", {type: "success"});
        e.preventDefault()
    }
    return(
        <div className="page_login">
            <div className="login_header">
                <img onClick = {() => navigate("/welcome")} src={arrowIcon} alt="Um icone preto de uma seta apontada para a esquerda"/>
                <h2>Cadastrar</h2>
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
                <button onClick={cadastrar} className="btn btn_input">Cadastrar</button>
            </form>
            
            <div className="form_division">
                <hr />
                <p>ou</p>
                <hr />
            </div>

            <div className="login_options">
                <button className="btn login_option_google"><img src={googleIcon}/><p>Cadastrar com google</p></button>
                <button className="btn login_option_linkedin"><img src={linkedinIcon}/><p>Cadastrar com linkedin</p></button>

            </div>
            <ToastContainer limit={1}/>

        </div>
    )

}