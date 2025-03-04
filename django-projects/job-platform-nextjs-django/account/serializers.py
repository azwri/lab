from django.contrib.auth import get_user_model
from rest_framework import serializers

User = get_user_model()

class SignUpSerializer(serializers.ModelSerializer):
    class Meta:
        model = User
        fields = ['first_name', 'last_name', 'email', 'password']
        extra_kwargs = {
            'first_name': {'required': True, 'allow_blank': False},
            'last_name': {'required': True, 'allow_blank': False},
            'email': {'required': True, 'allow_blank': False},
            'password': {'required': True, 'allow_blank': False, 'min_length': 6}
        }

class UserSerializer(serializers.ModelSerializer):
    mobile = serializers.CharField(source='profile.mobile')
    class Meta:
        model = User
        fields = ['first_name', 'last_name', 'email', 'username', 'mobile']