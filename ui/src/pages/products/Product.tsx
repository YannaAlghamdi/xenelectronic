import './Product.css';
import HeaderContainer from '../../components/header/HeaderContainer';
import { IonContent, IonHeader, IonPage, IonTitle, IonToolbar, IonCard, IonCardHeader, IonCardSubtitle, IonCardTitle, IonCardContent, IonItem, IonIcon, IonLabel, IonButton, IonRow, IonCol, IonButtons, IonImg, IonThumbnail } from '@ionic/react';
import { RouteComponentProps, withRouter } from 'react-router';
import axios from 'axios';
import React from 'react';
import { useToast } from '@agney/ir-toast';


interface ProductPageProps extends RouteComponentProps<{
    id: string;
    category: string
  }> {}

const  endpoint  =  `https://xenelectronic-app.herokuapp.com`;


const Product: React.FC<ProductPageProps> = ({match}) => {
   
  const Toast = useToast();
    const addToCart = (productId: any) => {
      axios.put(`${endpoint}/carts`, { 
          "product_id": productId,
          "cart_id": localStorage.getItem('cartId'),
      })
        .then(res => {
          const toast = Toast.create({ message: 'Product successfully added to cart', duration: 2000 });
          toast.present();
        })
      }
    
    const sendGetRequest = () => {
      return axios({
      method: 'get',
          url: `${endpoint}/categories/${match.params.id}/products`,
          responseType: 'stream'
      }).then(response => {
          console.log(response);
          return response.data;
      })
    };
    const [items, setItems] = React.useState([]);

    React.useEffect(() => {
      sendGetRequest().then(data => setItems(data));
    }, []);
    return (
      <IonPage>
      <HeaderContainer/>

        <IonContent>
          <IonRow className="row">
            <IonCol  className="ion-text-center category-label">
              <IonLabel>{match.params.category}</IonLabel>
            </IonCol>
          </IonRow>
          <IonRow className="row">
          {
            items.map((item, index) => {
              return (
                
                  <IonCol size="6" key={item['id']}>
                    <IonCard className="card"
                    routerDirection="forward">
                      <IonCardContent className="ion-text-center">
                        <IonImg className="item-img" src={item['photo']} />
                      <IonCardTitle>{item['name']}</IonCardTitle>
                      <IonCardSubtitle className="subtitle">{item['description']}</IonCardSubtitle>
                      <IonButton onClick={() => addToCart(item['id'])} color="light">Add to Cart</IonButton>
                      </IonCardContent>
                    </IonCard>
                  </IonCol>
              );
            })
          }
          </IonRow>
        </IonContent>
      </IonPage>
    );
};

export default withRouter(Product);
