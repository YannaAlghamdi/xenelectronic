import { IonContent, IonHeader, IonPage, IonTitle, IonToolbar, IonCard, IonCardHeader, IonCardSubtitle, IonCardTitle, IonCardContent, IonItem, IonIcon, IonLabel, IonButton, IonRow, IonCol, IonButtons, IonImg, IonThumbnail, IonInput, IonSelect, IonSelectOption } from '@ionic/react';
import './Checkout.css';
import { personCircle, search, helpCircle, star, create, ellipsisHorizontal, ellipsisVertical, cartOutline, trashOutline } from 'ionicons/icons';
import HeaderContainer from '../../components/header/HeaderContainer';
import axios from 'axios';
import React from 'react';
import { useToast } from '@agney/ir-toast';
import { useHistory } from 'react-router';

interface ContainerProps { }

const  endpoint  =  `http://localhost:8080`;

const Checkout: React.FC<ContainerProps> = () => {

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
  const [paymentMethod, setPaymentMethod] = React.useState<string>();
  const [accountName, setAccountName] = React.useState<string>();
  const [emailAddress, setEmailAddress] = React.useState<string>();
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
  const Toast = useToast();
  const history = useHistory();

  const checkout = () => {
    axios.post(`${endpoint}/orders`, { 
        "accountName": accountName,
        "emailAddress": emailAddress,
        "paymentMethod": paymentMethod,
        "cart": {
          "id": localStorage.getItem('cartId')
        }
    })
      .then(res => {
        history.push(`/checkout/success/${res.data.id}`);
      })
    }
  
  return (
    <IonPage>
      <HeaderContainer/>
      <IonContent>
        <IonCard className="checkout-card">
          <IonCardTitle className="ion-text-center card-title">Checkout</IonCardTitle>
          <IonCardContent>
            <form className="ion-padding">
              <IonRow className="row">
              <IonCol>

                  <IonRow>
                    <IonCol size="5">

                    </IonCol>
                    <IonCol size="4" className="header">
                      Item
                    </IonCol>
                    <IonCol size="3" className="header">
                      Price
                    </IonCol>
                  </IonRow>
                  
                  {
                    items.length ? items.map((item, index) => {
                    return (
                        <IonRow key={item['id']}>
                          <IonCol size="5">
                              <IonImg className="item-thumbnail" src={item['photo']} />
                          </IonCol>
                          <IonCol size="4" className="col-item">
                            {item['name']}
                          </IonCol>
                          <IonCol size="3" className="col-item">
                            {item['price']}
                          </IonCol>
                        
                        </IonRow>
                    );
                  }) : 
              <IonRow>
                <IonCol className="ion-text-center no-item-col">
                  <IonLabel className="no-item">There are no items in your cart.</IonLabel>
                </IonCol>
              </IonRow>
            }
                <IonRow>
                  <IonCol size="5">
                  </IonCol>
                  <IonCol size="4" className="col-item header ion-text-center">
                    Total
                  </IonCol>
                  <IonCol size="3" className="col-item header">
                    {getTotal().toFixed(2)}
                  </IonCol>
                </IonRow> 
                
              </IonCol>
              </IonRow>
              <IonItem>
                <IonLabel position="stacked" >Full Name</IonLabel>
                <IonInput value={accountName} onIonChange={e => setAccountName((e.target as HTMLInputElement).value)}/>
              </IonItem>
              <IonItem>
                <IonLabel position="stacked">Email Address</IonLabel>
                <IonInput  value={emailAddress} onIonChange={e => setEmailAddress((e.target as HTMLInputElement).value)} />
              </IonItem>
              <IonItem>
                <IonLabel position="stacked">Payment Method</IonLabel>
                <IonSelect value={paymentMethod} placeholder="Select One" onIonChange={e => setPaymentMethod(e.detail.value)}>
                <IonSelectOption value="Credit Card">Credit Card</IonSelectOption>
                <IonSelectOption value="Debit Card">Debit Card</IonSelectOption>
                <IonSelectOption value="Online Banking">Online Banking</IonSelectOption>
                <IonSelectOption value="OTC">Over-the-counter</IonSelectOption>

              </IonSelect>
              </IonItem>
              <IonItem>
                <IonLabel position="stacked">Amount</IonLabel>
                <IonInput value={getTotal().toFixed(2)} readonly/>
              </IonItem>

            </form>

            <IonButton onClick={() => checkout()} className="ion-margin-top btn-submit" type="submit" expand="block">
                  Place Order
              </IonButton>
          </IonCardContent>
        </IonCard>
       
      </IonContent>
    </IonPage>
  );
};

export default Checkout;
