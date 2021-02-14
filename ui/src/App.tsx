import { Redirect, Route, withRouter } from 'react-router-dom';
import { IonApp, IonRouterOutlet } from '@ionic/react';
import { IonReactRouter } from '@ionic/react-router';
import Home from './pages/Home';

/* Core CSS required for Ionic components to work properly */
import '@ionic/react/css/core.css';

/* Basic CSS for apps built with Ionic */
import '@ionic/react/css/normalize.css';
import '@ionic/react/css/structure.css';
import '@ionic/react/css/typography.css';

/* Optional CSS utils that can be commented out */
import '@ionic/react/css/padding.css';
import '@ionic/react/css/float-elements.css';
import '@ionic/react/css/text-alignment.css';
import '@ionic/react/css/text-transformation.css';
import '@ionic/react/css/flex-utils.css';
import '@ionic/react/css/display.css';

/* Theme variables */
import './theme/variables.css';
import Product from './pages/products/Product';
import { ToastProvider } from '@agney/ir-toast';
import ShoppingCart from './pages/shopping-cart/ShoppingCart';

const App: React.FC = () => (
  <IonApp>
    <ToastProvider>
      <IonReactRouter>
        <IonRouterOutlet>
          <Route exact path="/home">
            <Home />
          </Route>
          <Route exact path="/">
            <Redirect to="/home" />
          </Route>
          <Route exact path="/:id/:category/products">
            <Product/>
          </Route>
          <Route exact path="/cart">
            <ShoppingCart/>
          </Route>
        </IonRouterOutlet>
      </IonReactRouter>
    </ToastProvider>
  </IonApp>
);

export default App;
