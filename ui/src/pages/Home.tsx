import { IonBackButton, IonButton, IonButtons, IonContent, IonHeader, IonIcon, IonLabel, IonPage, IonTitle, IonToolbar } from '@ionic/react';
import HomeContainer from '../components/home/HomeContainer';
import './Home.css';
import { personCircle, search, helpCircle, star, create, ellipsisHorizontal, ellipsisVertical, cartOutline } from 'ionicons/icons';
import HeaderContainer from '../components/header/HeaderContainer';
import axios from 'axios';


const  endpoint  =  `http://localhost:8080`;

const createCart = () => {
  axios.post(`${endpoint}/carts`, { })
    .then(res => {
      localStorage.setItem('cartId', res.data.id);
    })
  }

const Home: React.FC = () => {
  createCart();
  return (
    <IonPage>
      <HeaderContainer/>
        <HomeContainer />
    </IonPage>
  );
};

export default Home;
