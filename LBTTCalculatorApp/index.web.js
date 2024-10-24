import { AppRegistry } from 'react-native';
import App from './App';
import { name as appName } from './app.json';
import { render } from 'react-dom';

AppRegistry.registerComponent(appName, () => App);

const rootTag = document.getElementById('app-root');
AppRegistry.runApplication(appName, { initialProps: {}, rootTag });
