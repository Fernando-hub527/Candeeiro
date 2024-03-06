import personIcon from "../../static/img/personIcon.png"
import "../../static/css/stylePresentation.css"
import { useNavigate } from "react-router-dom";

export function Presentation(props){
    const navigate = useNavigate();
    
    console.log(props)
    const openScreen = (path) => setTimeout(()=>navigate(path), 150)

    return <div className="page_presentation">
        <header>
            <h1>Bem-vindo</h1>
            <p>Essa é uma ferramenta de baixo custo para acompanhamento do consumo de energia da sua casa</p>
        </header>
        <img className="page_presentation_user" src={personIcon} alt="icone de uma pessoa vestindo blusa roxa e short amarelo enquanto equilibra equipamentos eletrônicos"/>
       
        <div className="form_btn">
            <button className="btn btn-login" onClick={() => openScreen("/login")}>LOG IN</button>
            <button className="btn btn-register">REGISTER</button>
        </div>

        <p className="page_presentation_obs">Esse projeto faz parte do meu <a href="">portifólio</a> e possui carater apenas demonstrativo</p>
    </div>
}