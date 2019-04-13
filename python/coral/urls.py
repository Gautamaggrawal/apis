from django.views.generic import TemplateView
from django.contrib import admin
from django.urls import include, path
from rest_framework import routers
from api import views



router = routers.DefaultRouter()
router.register(r'users', views.UserViewSet)

# Wire up our API using automatic URL routing.
# Additionally, we include login URLs for the browsable API.

urlpatterns = [
    path(
        'admin/',
        admin.site.urls),
    path(
        'api/',
        include(
            router.urls)),
    path('', TemplateView.as_view(template_name='home.html'))
 ]
