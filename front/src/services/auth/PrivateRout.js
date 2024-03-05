import cookie from "js-cookies"

export const PrivateRoute = ({ children }) => {
    
    if (!cookie.getItem("access_token")) document.location.replace("/presentation") 
    return children
};

