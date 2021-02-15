import { IonButton, IonButtons, IonHeader, IonIcon, IonTitle, IonToolbar } from '@ionic/react';
import './HeaderContainer.css';
import { ellipsisHorizontal, ellipsisVertical, cartOutline } from 'ionicons/icons';
import { Link } from 'react-router-dom';
import React from 'react';

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
