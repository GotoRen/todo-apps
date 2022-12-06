import React from "react";
import styled from "styled-components";

const FooterComponent = styled.footer`
  width: 100%;
  color: #fff;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.2);
  text-align: center;
  padding: 1px 0;
  position: absolute;
  bottom: 0;
`;

const FooterText = styled.a`
  color: #787878;
  text-align: center;
  font-size: 12px;
`;

const Footer = () => {
  return (
    <FooterComponent>
      <FooterText href="https://github.com/GotoRen/todo-apps">
        &copy; {new Date().getFullYear()} Ren Goto. All rights reserved. <br />
      </FooterText>
      <FooterText href="https://reactjs.org/">
        {` 🔨 `} Powered by React.js ver 18.2
      </FooterText>
    </FooterComponent>
  );
};

export default Footer;
