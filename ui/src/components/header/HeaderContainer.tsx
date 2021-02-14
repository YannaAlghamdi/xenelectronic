import { IonBackButton, IonButton, IonButtons, IonContent, IonHeader, IonIcon, IonLabel, IonPage, IonTitle, IonToolbar } from '@ionic/react';
import './HeaderContainer.css';
import { personCircle, search, helpCircle, star, create, ellipsisHorizontal, ellipsisVertical, cartOutline } from 'ionicons/icons';
import { Link } from 'react-router-dom';


const HeaderContainer: React.FC = () => {
  return (
      <IonHeader  className="ion-no-border">
        <IonToolbar>
          <IonButtons slot="secondary">
            <IonButton>
            </IonButton>
            <IonButton routerLink={"/cart"}
                 routerDirection="forward" >
              <IonIcon slot="icon-only"icon={cartOutline} />
            </IonButton>
          </IonButtons>
          <IonButtons slot="primary">
            <IonButton color="secondary">
              <IonIcon slot="icon-only" ios={ellipsisHorizontal} md={ellipsisVertical} />
            </IonButton>
          </IonButtons>
          <IonTitle><Link to="/home">XenElectronic</Link></IonTitle>
        </IonToolbar>
      </IonHeader>
  );
};

export default HeaderContainer;
