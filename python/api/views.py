from api.models import Userdata
from rest_framework import viewsets
from api.serializers import UserSerializer
from django_filters import rest_framework as filters
from rest_framework import filters


class UserViewSet(viewsets.ModelViewSet):
    """
    API endpoint that allows users to be viewed or edited.
    """
    queryset = Userdata.objects.all()
    serializer_class = UserSerializer
    filter_backends = (filters.SearchFilter,)
    search_fields = ('emailid',)
    # filter_backends = (filters.DjangoFilterBackend,)
    # filter_fields = ('emailid','username')
