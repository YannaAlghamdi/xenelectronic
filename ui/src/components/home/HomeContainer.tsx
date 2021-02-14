import { IonContent, IonHeader, IonPage, IonTitle, IonToolbar, IonCard, IonCardHeader, IonCardSubtitle, IonCardTitle, IonCardContent, IonItem, IonIcon, IonLabel, IonButton, IonRow, IonCol } from '@ionic/react';
import './HomeContainer.css';
import axios from 'axios';
import React from 'react';
import { laptopOutline, tvOutline, watchOutline, tabletPortraitOutline, phonePortraitOutline } from 'ionicons/icons';

interface ContainerProps { }

const  endpoint  =  `https://xenelectronic-app.herokuapp.com`;

const sendGetRequest = () => {
	return axios({
    method: 'get',
		url: `${endpoint}/categories`,
		responseType: 'stream'
  }).then(response => {

    console.log(response);
    return response.data;
  })
};


const HomeContainer: React.FC<ContainerProps> = () => {
	const [items, setItems] = React.useState([]);
  const itemIcons = [laptopOutline, tvOutline, watchOutline, tabletPortraitOutline, phonePortraitOutline]
  console.log(itemIcons);
	React.useEffect(() => {
		sendGetRequest().then(data => setItems(data));
	}, []);
  return (
		<IonContent>
      <IonRow className="row">
			{
        items.map((item, index) => {
          return (
            
              <IonCol size="6" key={item['id']}>
                <IonCard
                 routerLink={`/${item['id']}/${item['name']}/products`}
                 routerDirection="forward">
                  <IonCardContent className="ion-text-center">
                  <IonIcon className="card-icon" slot="icon-only" icon={itemIcons[index]}/>
                  <IonCardTitle>{item['name']}</IonCardTitle>
                  </IonCardContent>
                </IonCard>
              </IonCol>
          );
        })
      }
			</IonRow>
		</IonContent>
  );
};

export default HomeContainer;
