import {Typography} from 'antd';
import styled, {css} from 'styled-components';

export const Header = styled.div<{visiblePortion: number}>`
  display: flex;
  align-items: center;
  align-content: center;
  cursor: pointer;
  width: 100%;
  background: #F5F5FA;
  ${props =>
    css`
      height: ${props.visiblePortion}px;
    `}
  padding: 0 24px;
  color: rgb(213, 215, 224);
`;

export const HeaderText = styled(Typography.Text)``;

export const StartDateText = styled(Typography.Text)`
  && {
    margin-left: 14px;
    margin-right: 40px;
    font-size: 12px;
  }
`;

export const CountNumber = styled.span`
  margin-right: 15px;
`;

export const Container = styled.div`
  border: 1px solid rgba(0, 0, 0, 0.06);
  margin-bottom: 16px;
  overflow-y: hidden;
`;

export const Content = styled.div`
  padding: 24px;
  overflow-y: scroll;
  height: 330px;
`;