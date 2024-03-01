import "../../static/css/styleLoginPage.css"
import logo from "../../static/img/logo.svg"
import "@fontsource/comfortaa"

export function LoginPage(props) {
    return <div className="page_login">
        <div className="login_header">
            <img src={logo} />
            <h1>Candieiro</h1>
        </div>
        <form>
            <div className="form_input form_user">
                <input type="text" id="nickName" placeholder="Usuário" />
            </div>
            <div className="form_input form_password">
                <input type="password" id="password" placeholder="Senha" />
            </div>
            <div className="form_btn">
                <button className="btn btn-input">LOG IN</button>
                <button className="btn btn-register">REGISTER</button>
            </div>
        </form>
    </div>


}