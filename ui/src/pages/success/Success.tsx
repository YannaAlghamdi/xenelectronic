import { IonCard, IonCardContent, IonCol, IonContent, IonLabel, IonPage, IonRow } from '@ionic/react';
import './Success.css';
import HeaderContainer from '../../components/header/HeaderContainer';
import React from 'react';
import { Link } from 'react-router-dom';

interface ContainerProps { }

const Success: React.FC<ContainerProps> = () => {

  return (
    <IonPage>
      <HeaderContainer/>
      <IonContent>
          <IonRow>
              <IonCol>
                <IonCard className="success-card">
                  <IonCardContent className="ion-text-center">
                  <IonLabel className="success-label">Thank you for shopping with us!</IonLabel>
                  <small><Link to={"/home"}>BACK TO HOME</Link></small>
                  </IonCardContent>
                </IonCard>
              </IonCol>
          </IonRow>
      </IonContent>
    </IonPage>
  );
};

export default Success;
