import personIcon from "../../static/img/personIcon.png"
import "../../static/css/stylePresentation.css"

export function Presentation(props){
    return <div className="page_presentation">
        <header>
            <h1>Bem-vindo</h1>
            <p>Essa é uma ferramenta de baixo custo para acompanhamento do consumo de energia da sua casa</p>
        </header>
        <img className="page_presentation_user" src={personIcon} alt="icone de uma pessoa vestindo blusa roxa e short amarelo enquanto equilibra equipamentos eletrônicos"/>
       
        <div className="form_btn">
            <button className="btn btn-login">LOG IN</button>
            <button className="btn btn-register">REGISTER</button>
        </div>

        <p className="page_presentation_obs">Esse projeto faz parte do meu <a href="">portifólio</a> e possui carater apenas demonstrativo</p>
    </div>
}