import { IonContent, IonHeader, IonPage, IonTitle, IonToolbar, IonCard, IonCardHeader, IonCardSubtitle, IonCardTitle, IonCardContent, IonItem, IonIcon, IonLabel, IonButton, IonRow, IonCol, IonButtons, IonImg, IonThumbnail } from '@ionic/react';
import './ShoppingCart.css';
import { personCircle, search, helpCircle, star, create, ellipsisHorizontal, ellipsisVertical, cartOutline, trashOutline } from 'ionicons/icons';
import HeaderContainer from '../../components/header/HeaderContainer';
import axios from 'axios';
import React from 'react';

interface ContainerProps { }

const  endpoint  =  `http://localhost:8080`;

const ShoppingCart: React.FC<ContainerProps> = () => {

  const getShoppingCart = () => {
    return axios({
      method: 'get',
        url: `${endpoint}/carts/${localStorage.getItem('cartId')}/products`,
        responseType: 'stream'
    }).then(response => {
      console.log(response);
      return response.data;
    })
  };

  const [items, setItems] = React.useState([]);
    React.useEffect(() => {
      getShoppingCart().then(data => setItems(data));
  }, []);
  
  const getTotal = () => {
    var sum = 0;
    items.forEach(item => {
      sum = sum + parseInt(item['price']);
    })
    return sum;
  }

  return (
    <IonPage>
      <HeaderContainer/>
      <IonContent>
      <IonRow className="row">
        <IonCol>
          <IonCard className="card">
            <IonCardTitle className="ion-text-center">Your Shopping Cart</IonCardTitle>
            <IonCardContent>
              <IonRow>
                <IonCol size="5">

                </IonCol>
                <IonCol size="3" className="header">
                  Item
                </IonCol>
                <IonCol size="2" className="header">
                  Price
                </IonCol>
                <IonCol size="2">
                  
                </IonCol>
              </IonRow>
              {
            items.map((item, index) => {
              return (
                  <IonRow key={item['id']}>
                    <IonCol size="5">
                        <IonImg className="item-thumbnail" src={item['photo']} />
                    </IonCol>
                    <IonCol size="3" className="col-item">
                      {item['name']}
                    </IonCol>
                    <IonCol size="2" className="col-item">
                      {item['price']}
                    </IonCol>
                    <IonCol size="2">
                      <IonButton color="danger" class="col-remove-item">
                        <IonIcon slot="icon-only" icon={trashOutline}></IonIcon>
                      </IonButton>
                    </IonCol>
                  </IonRow>
              );
            })
          }
            <IonRow>
              <IonCol size="5">
              </IonCol>
              <IonCol size="3" className="col-item header ion-text-center">
                Total
              </IonCol>
              <IonCol size="2" className="col-item header">
                 {getTotal().toFixed(2)}
              </IonCol>
              <IonCol size="2">
                  
              </IonCol>
            </IonRow>   
            </IonCardContent>
          </IonCard>
        </IonCol>
          {/* {
            items.map((item, index) => {
              return (
                  <IonCol size="6" key={item['id']}>
                    <IonCard className="card"
                    routerDirection="forward">
                      <IonCardContent className="ion-text-center">
                      <IonCardTitle>{item['name']}</IonCardTitle>
                      </IonCardContent>
                    </IonCard>
                  </IonCol>
              );
            })
          } */}
          </IonRow>
        </IonContent>
    </IonPage>
  );
};

export default ShoppingCart;
