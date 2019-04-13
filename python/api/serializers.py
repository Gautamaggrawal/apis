from api.models import Userdata
from rest_framework import serializers
from rest_framework.validators import UniqueValidator
from django.contrib.auth.hashers import make_password


class UserSerializer(serializers.HyperlinkedModelSerializer):
    password = serializers.CharField(
        write_only=True,
        required=True,
        help_text='Leave empty if no change needed',
        style={'input_type': 'password', 'placeholder': 'Password'})

    emailid = serializers.EmailField()

    def validate_phoneno(self, data):
        if len(data) == 10 and data.isalnum() and(data.startswith(
                "6") or data.startswith("7") or data.startswith("8") or data.startswith("9")):
            return data
        else:
            raise serializers.ValidationError("incorrect phone number")

    class Meta:
        model = Userdata
        fields = ('username', 'password', 'emailid', 'phoneno',)

    def create(self, validated_data):
        validated_data['password'] = make_password(
            validated_data.get('password'))
        user, created = Userdata.objects.update_or_create(
            emailid=validated_data.get(
                'emailid', None), defaults={
                'password': validated_data.get(
                    'password', None), "phoneno": validated_data.get(
                    'phoneno', None), "username": validated_data.get(
                        'username', None)})
        return user
        # return super(UserSerializer,self).create(validated_data)
