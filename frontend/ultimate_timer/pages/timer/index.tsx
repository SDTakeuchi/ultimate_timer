import { ListItem } from '@mui/material'

interface PropsType {
    children: React.ReactNode;
  }
  
  export default function ListView(props: PropsType) {
    return (
      <ul>
        {React.Children.map(props.children, (child) => {
          return <li>{child}</li>;
        })}
      </ul>
    );
  }
