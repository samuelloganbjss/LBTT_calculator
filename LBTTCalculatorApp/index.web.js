import { AppRegistry } from 'react-native';
import App from './App';
import appConfig from './app.json'; 
import { render } from 'react-dom';

AppRegistry.registerComponent(appConfig.name, () => App); 

const rootTag = document.getElementById('app-root');
AppRegistry.runApplication(appConfig.name, { initialProps: {}, rootTag });
