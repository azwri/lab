from django.urls import path
from .views import register, get_current_user
from rest_framework_simplejwt.views import TokenObtainPairView, TokenVerifyView

app_name = 'account'

urlpatterns = [
    path('register/', register, name='register'),
    path('user/', get_current_user, name='user'),
    path('token/', TokenObtainPairView.as_view(), name='token_obtain_pair'),
    path('token/verify/', TokenVerifyView.as_view(), name='token_verify'),
]