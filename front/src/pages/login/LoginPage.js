import "../../static/css/styleLoginPage.css"
import "@fontsource/comfortaa"
import arrowIcon from "../../static/img/arrowIcon.svg"
import eyeIcon from "../../static/img/eyeIcon.svg"
import { useState } from "react"

export function LoginPage(props) {
    const [passwordVisibility, setVisibility] = useState(false)


    return <div className="page_login">
        <div className="login_header">
            <img src={arrowIcon} alt="Um icone preto de uma seta apontada para a esquerda"/>
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
            <button className="btn btn_input">Entrar</button>
        </form>
        <div className="form_division">
            <hr />
            <p>ou</p>
            <hr />
        </div>

        <div className="login_options">
            <button className="btn login_option">Continuar com google</button>
            <button className="btn login_option"><img src={eyeIcon}/>Continuar com linkedin</button>

        </div>
    </div>


}