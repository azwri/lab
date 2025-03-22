from django.urls import path, include
from . import views
from rest_framework import routers

app_name = 'api'
router = routers.DefaultRouter()
router.register(r'movies', views.MovieViewSet)
router.register(r'ratings', views.RatingViewSet)

urlpatterns = [
    path('', include(router.urls)),
]
